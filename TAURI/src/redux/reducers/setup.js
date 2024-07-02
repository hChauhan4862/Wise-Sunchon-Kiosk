
const setupReducer = (state = { selectedLanguage: 'ENGLISH' }, action) => {
    switch (action.type) {
        case 'SET_SERVER_SYNC':
            return { ...state, hcLiveSync: action.data }


        case 'SET_KIOSK_CONFIG':
            return { ...state, kioskConfig: action.data }
        case 'SET_TIMERS':
            return { ...state, timers: action.data }
        case 'SET_NOTICE':
            return { ...state, notices: action.data }
        case 'SET_SELECTED_LANGUAGE':
            return { ...state, selectedLanguage: action.data }
        default:
            return state;
    }
}

export default setupReducer