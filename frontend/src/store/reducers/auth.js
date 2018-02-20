import * as actionTypes from '../actions/actionTypes';
import { updateObject } from "../../shared/utility";
import {cleanError} from "../actions";

const initialState = {
    authKey: null,
    error: null,
    loading: false
};

const authStart = ( state, action ) => {
    return updateObject(state, {
        error: null,
        loading: true
    });
};

const authSuccess = ( state, action ) => {
    return updateObject(state, {
        authKey: action.authData.accessToken,
        error: null,
        loading: false
    });
};

const authFail = ( state, action ) => {
    return updateObject(state, {
        error: action.error,
        loading: false
    });
};

const authLogout = ( state, action ) => {
    return updateObject(state, {
        authKey: null,
        error: null,
        loading: false
    });
};

const authCleanError = (state, action) => {
    return updateObject(state, {
        error: null
    });
};

const reducer = (state = initialState, action) => {
    if (action) {
        switch (action.type) {
            case actionTypes.AUTH_START: return authStart(state, action);
            case actionTypes.AUTH_SUCCESS: return authSuccess(state, action);
            case actionTypes.AUTH_FAIL: return authFail(state, action);
            case actionTypes.AUTH_LOGOUT: return authLogout(state, action);
            case actionTypes.AUTH_CLEAN_ERROR: return authCleanError(state, action);
            default:
                return state;
        }
    }

    return state;
};

export default reducer;