import React, {Fragment, useEffect, useState} from "react";

import Title from "../Title";
import Ticket from "../Ticket";
import Fab from "../Fab";
import {useDispatch, useSelector} from "react-redux";
import {createProject, listProjects} from "../../actions/api";
import {openProject} from "../../actions/navigation";
import Form, {PROJECT_FORM_CONFIGURATION} from "../Form";

const ProjectContainer = () => {
    const [addProject, toggleProjectForm] = useState(false);
    const projects = useSelector(state => state.projects);
    const dispatch = useDispatch();

    useEffect(() => {
        dispatch(listProjects())
    }, [])

    const projectList = projects.map(item => <Ticket key={item.id} title={item.title}
                                                     description={item.description}
                                                     action={() => dispatch(openProject(item.id))}/>
    );

    if (addProject) {
        function submitHandler(title, description) {
            dispatch(createProject(...arguments))
            toggleProjectForm(!addProject)
        }

        return <Form config={PROJECT_FORM_CONFIGURATION} submit={submitHandler} close={() => toggleProjectForm(false)} />
    }

    return <Fragment>
        <Title title="My Projects" />

        { projectList }

        <Fab onClick={() => toggleProjectForm(!addProject)} />
    </Fragment>;
}

export default ProjectContainer;