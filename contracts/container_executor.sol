contract token_bank {
    function create_escrow(address _cp, address _contract, uint _amount) returns (bool rv) {}
    function cancel_escrow(address _proposer, address _cp, address _contract, uint _amount) returns (bool rv) {}
}

contract container_executor {
    address owner;
    string whisper;
    string agreement;
    address container_provider;
    token_bank piggy_bank;

    uint constant new_container_event_code = 1;
    uint constant execution_complete_event_code = 2;
    uint constant container_rejected_event_code = 3;
    uint constant container_cancelled_event_code = 4;

    event NewContainer(uint indexed _eventcode, string _id, address indexed _self);
    event ExecutionComplete(uint indexed _eventcode, string _id, address indexed _self);
    event ContainerRejected(uint indexed _eventcode, string _id, address indexed _self);
    event ContainerCancelled(uint indexed _eventcode, string _id, address indexed _self);

    // This function is used by governors (Glensung, IoT providers, etc) to make proposals to
    // device owners for running containers. This function escrows the offered funds in the
    // token bank.
    function new_container(string _whisperId, string _agreementId, uint _amount) returns (bool r) {
        if (_amount == 0) {
            return false;
        }
        if (!in_contract()) {
            var escrowed = piggy_bank.create_escrow(owner,address(this),_amount);
            if (escrowed == true) {
                whisper = _whisperId;
                agreement = _agreementId;
                container_provider = tx.origin;
                NewContainer(new_container_event_code,_agreementId,this);
                return true;
            } else {
                return false;
            }
        } else {
            return false;
        }
    }

    // This function is used by device owners to reject proposals made by governors (Glensung,
    // IoT providers, etc). The proposal will be cancelled and any funds will be distributed
    // when the governer closes out the agreement.
    function reject_container() returns (bool r) {
        if (tx.origin == owner) {
            piggy_bank.cancel_escrow(container_provider, tx.origin, this, 0);
            ContainerRejected(container_rejected_event_code,agreement,this);
            clear_container();
            return true;
        } else {
            return false;
        }
    }

    // This function is used by governors (Glensung, IoT providers, etc) to cancel agreements with
    // device owners for running containers. Input amount funds are transferred to the device owner,
    // remaining escrowed funds are returned to the proposer.
    function cancel_container(uint _amount) returns (bool r) {
        if (container_provider != address(0) && tx.origin == container_provider) {
            piggy_bank.cancel_escrow(tx.origin, owner, this, _amount);
            ContainerCancelled(container_rejected_event_code,agreement,this);
            clear_container();
            return true;
        } else {
            return false;
        }
    }

    // This function is internal, used to clear out agreement state.
    function clear_container() internal returns (bool r) {
        whisper = "";
        container_provider = address(0);
        agreement = "";
    }

    // This function is used to determine agreement status. Is there an agreement in place
    // with another party or not.
    function in_contract() constant returns (bool r) {
        if (container_provider == address(0)) {
            return false;
        } else {
            return true;
        }
    }

    // Function to retrieve agreement state from the system.
    function get_agreement_id() constant returns (string r) {
        return agreement;
    }
    function get_container_provider() constant returns (address r) {
        return container_provider;
    }
    function get_whisper() constant returns (string r) {
        return whisper;
    }

    // Constructor and other infrastructure functions
    function container_executor() {
        owner = msg.sender;
    }
    function set_bank(address _bank) {
        if (owner == msg.sender) {
            piggy_bank = token_bank(_bank);
        }
    }
    function get_bank() constant returns (address r) {
        return piggy_bank;
    }
    function get_owner() constant returns (address r) {
        return owner;
    }

    // Used to get rid of the contract
    function kill() {
        if (msg.sender == owner) suicide(owner);
    }
}

