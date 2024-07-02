
const authReducer = (state = {}, action) => {
    switch (action.type) {
        case 'SET_USER_DATA':
            return { ...state, userData: action.data }
        case 'SET_USER_ID':
            return { ...state, userId: action.user }
        case 'GROUP_STUDY_MEMBERS':
            return { ...state, members: action.members}
        default:
            return state;
    }
}

export default authReducer