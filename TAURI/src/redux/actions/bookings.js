import * as api from '../../apis'

export const printTickit = (data, successAlert, errorAlert) => async () => {
    console.log(data, 'from print')
    try {
        const res = await api.printTickit(data)
        successAlert()
    } catch (error) {
        // errorAlert()
        console.log(error)
    }
}