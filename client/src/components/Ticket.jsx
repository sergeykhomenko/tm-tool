import React from "react";

const Ticket = ({title, description, action}) => {
    return <div className="ticket">
        <div className="priority">
            <i className="priority__mark priority__mark--high">&nbsp;</i>
        </div>
        <div className="complete complete--unchecked">&nbsp;</div>
        <div className="ticket-info" onClick={ action }>
            <span className="ticket__title">{ title }</span>
            <span className="ticket__tags">{ description }</span>
        </div>
    </div>
};

export default Ticket;