import React from 'react';

export const Title = ({label, light = false}) => {
    let classNames = ['title'];
    if (light) {
        classNames.push('title--light')
    }

    return <h1 className={classNames.join(' ')}>{label}</h1>
}


export const Text = ({ children }) => {
    return <span>{ children }</span>
}