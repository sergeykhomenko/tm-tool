import React from 'react';
import ReactDOM from 'react-dom';

import './index.css';
import './App.scss';

import App from './App';
import * as serviceWorker from './serviceWorker';
import {applyMiddleware, compose, createStore} from "redux";
import createSagaMiddleware from 'redux-saga';
import {initialState, rootReducer} from "./reducers/rootReducer";
import axios from 'axios';
import axiosMiddleware from "redux-axios-middleware";
import {Provider} from "react-redux";

const client = axios.create({
    baseURL: 'http://localhost:8080/api',
    responseType: 'json'
});
const sagaMiddleware = createSagaMiddleware();
const composeEnhancers = (typeof window !== 'undefined' && window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__) || compose;
const store = createStore(
    rootReducer,
    initialState,
    composeEnhancers(
        applyMiddleware(sagaMiddleware, axiosMiddleware(client)),
    ));

ReactDOM.render(
    <React.StrictMode>
        <div className="app">
            <Provider store={store}>
                <App/>
            </Provider>
        </div>
    </React.StrictMode>,
    document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
