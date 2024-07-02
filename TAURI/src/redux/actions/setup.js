import * as api from '../../apis'

export const setTimers = (timers) => async (dispatch) => {
    try {
        dispatch({ type: 'SET_TIMERS', data: timers})
    } catch (err) {
        console.log(err)
    }
}

export const hcLiveSyncFun = () => async (dispatch) => {
    try {
        const res = await api.getHcLiveSync()
        dispatch({ type: 'SET_SERVER_SYNC', data: res.data.result })
    } catch (error) {
        // dispatch({ type: 'SET_SERVER_SYNC', data: null })
        console.log(error)
    }
}

export const setSelectedLanguage = (language) => async (dispatch) => {
    dispatch({ type: 'SET_SELECTED_LANGUAGE', data: language })
}