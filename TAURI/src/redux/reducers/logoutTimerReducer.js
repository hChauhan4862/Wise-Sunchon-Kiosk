

const logoutTimerReducer = (state = {}, action) => {
    switch (action.type) {
        case 'SET_LOGOUT_TIMER':
            return { ...state, resetLogoutTimerDispatch: action.data }
        default:
            return state;
    }
}

export default logoutTimerReducer;