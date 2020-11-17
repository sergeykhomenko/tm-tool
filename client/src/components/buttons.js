import React from "react";

const ButtonComponent = (content = '', modifiers = []) => {
    return <div className={['btn', ...modifiers].join(' ')}>
        <span>{content}</span>
    </div>;
}

export const Button = () => ButtonComponent('Button');