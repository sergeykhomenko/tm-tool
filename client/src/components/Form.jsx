import React, {Fragment, useState} from "react";
import Title from "./Title";
import Back from "./Back";

export const PROJECT_FORM_CONFIGURATION = {
    title: 'Create new project',
    name: 'Project name',
    description: 'Project description'
};

export const TICKET_FORM_CONFIGURATION = {
    title: 'Create new task',
    name: 'Ticket title',
    description: 'Description'
};

const Form = ({config, close, submit}) => {
    const {title, name, description} = config;

    const [titleValue, setTitle] = useState('');
    const [descriptionValue, setDescription] = useState('');

    return <Fragment>
        <Back action={close} />
        <Title title={title} />

        <div className="form">
          <input type='text' className="form__content" placeholder={name} onChange={e => setTitle(e.target.value)} />
          <textarea className="form__content" placeholder={description} onChange={e => setDescription(e.target.value)} />
          <button className="button" onClick={() => submit(titleValue, descriptionValue)}>Submit</button>
        </div>
    </Fragment>
}

export default Form;