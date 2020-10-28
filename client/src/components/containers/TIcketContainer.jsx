import db from '../../db.json';
import {useDispatch, useSelector} from "react-redux";
import Title from "../Title";
import React, {Fragment, useState} from "react";
import Back from "../Back";
import Fab from "../Fab";
import {createProject, createTicket} from "../../actions/api";
import Form, {PROJECT_FORM_CONFIGURATION, TICKET_FORM_CONFIGURATION} from "../Form";

const TicketContainer = () => {
    const project = useSelector(state => {
        let projectIndex = state.projects.findIndex(item => item.id === state.currentProject);

        if (projectIndex === -1)
            return null;

        return state.projects[projectIndex];
    });

    const [addTicket, toggleTicketForm] = useState(false);
    const dispatch = useDispatch();

    if (project === null)
        return null;

    // const tickets = db.projects[currentProject].tickets.map((item, key) => <div className="ticket" key={key}>
    //     <div className="priority">
    //         <i className="priority__mark priority__mark--high">&nbsp;</i>
    //     </div>
    //     <div className="complete complete--unchecked">&nbsp;</div>
    //     <div className="ticket-info">
    //         <span className="ticket__title">{item.title}</span>
    //         <span className="ticket__tags">{item.description} <br/>@{item.author}</span>
    //     </div>
    // </div>);

    if (addTicket) {
        function submitHandler(title, description) {
            dispatch(createTicket(project.id, ...arguments))
            toggleTicketForm(!addTicket)
        }

        return <Form config={TICKET_FORM_CONFIGURATION} submit={submitHandler} close={() => toggleTicketForm(false)} />
    }

    return <Fragment>
        <Back />
        <Title title={project.title} />

        <div className="description">{project.description}</div>

        <Fab onClick={() => toggleTicketForm(!addTicket)} />

        {/*{tickets}*/}
    </Fragment>;
}

export default TicketContainer;