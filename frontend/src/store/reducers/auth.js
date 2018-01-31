import * as actionTypes from '../actions/actionTypes';
import { updateObject } from "../../shared/utility";

const initialState = {
    username: null,
    password: null,
    title: null,
    profileImage: null,
    authKey: null,
    patientTypeName: null,
    firstName: null,
    lastName: null,
    email: null,
    phone: null,
    phoneIsConfirmed: 0,
    birthday: null,
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
        username: action.username,
        password: action.password,
        title: action.authData.title,
        profileImage: action.authData.profile_image,
        authKey: action.authData.auth_key,
        patientTypeName: action.authData.patient_type_name,
        firstName:  action.authData.first_name,
        lastName: action.authData.last_name,
        email:  action.authData.email,
        phone: action.authData.phone,
        phoneIsConfirmed: action.authData.phone_is_confirmed,
        birthday: action.authData.birthday,
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

const authLogout = (state, action) => {
    return updateObject(state, {
        username: null,
        password: null,
        title: null,
        profileImage: null,
        authKey: null,
        patientTypeName: null,
        email:  null,
        phone: null,
        phoneIsConfirmed: 0,
        birthday: null,
        error: null
    });
};

const  updateUserInfoSuccess = (state, action) => {
    return updateObject(state, {
        title: action.authData.title,
        profileImage: action.authData.profile_image,
        patientTypeName: action.authData.patient_type_name,
        firstName:  action.authData.first_name,
        lastName: action.authData.last_name,
        email:  action.authData.email,
        phone: action.authData.phone,
        phoneIsConfirmed: action.authData.phone_is_confirmed,
        birthday: action.authData.birthday,
        error: null,
        loading: false
    });
};

const reducer = (state = initialState, action) => {
    if (action) {
        switch (action.type) {
            case actionTypes.AUTH_START: return authStart(state, action);
            case actionTypes.AUTH_SUCCESS: return authSuccess(state, action);
            case actionTypes.AUTH_FAIL: return authFail(state, action);
            case actionTypes.AUTH_LOGOUT: return authLogout(state, action);
            case actionTypes.UPDATE_USER_INFO_START: return authStart(state, action);
            case actionTypes.UPDATE_USER_INFO_SUCCESS: return updateUserInfoSuccess(state, action);
            case actionTypes.UPDATE_USER_INFO_FAIL: return authFail(state, action);
            default:
                return state;
        }
    }

    return state;
};

export default reducer;