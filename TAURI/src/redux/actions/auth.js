import * as api from '../../apis'
import { toast } from 'react-toastify';
import jwtDecode from 'jwt-decode'
import Swal from "sweetalert2";
import {hcAlert} from '../../components/Modal/hcAlert'


export const normalLogin = (loginData) => async (dispatch) => {
    try {
        const res = await api.getHcLogin(loginData)
        if (!res.data.result?.user){
            hcAlert('Opps!',res.data.result)
            return;
        }
        dispatch({ type: 'SET_USER_DATA', data: res.data.result })
        Swal?.close()
    } catch (err) {
        hcAlert('Opps!', err.response.data?.result || "Something went wrong")
    }
}

export const getUserByQr = (qrCode, toast) => async (dispatch) => {
    try {
        const res = await api.getHcLogin(qrCode)
        // if (!res.data.result.user){throw error}
        dispatch({ type: 'SET_USER_ID', user: res.data.result?.user?.SCHOOL_NO })
        return res.data.result?.user?.SCHOOL_NO
    } catch (error) {
        toast.error(error.response.data?.result || "Something went wrong")
    }
}

export const setUserForQR = (id) => async (dispatch) => {
    try {
        dispatch({ type: 'SET_USER_ID', user: id })
        return id
    } catch (error) {
        console.log(error)
    }
}

export const logoutUser = (navigate) => async (dispatch) => {
    dispatch({ type: 'SET_USER_DATA', data: null })
    if (navigate) {
        navigate('/')
    }
}

export const groupStudyMembers = (members) => async (dispatch) => {
    dispatch({ type: 'GROUP_STUDY_MEMBERS', members })
}
