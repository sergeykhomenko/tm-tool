import React from "react";
import {back} from "../actions/navigation";
import {useDispatch} from "react-redux";

const Back = ({action}) => {
    const dispatch = useDispatch();

    if (action !== undefined) {
        return <div className='back' onClick={action}>Back</div>;
    }

    return <div className='back' onClick={() => dispatch(back())}>Back</div>;
}

export default Back;