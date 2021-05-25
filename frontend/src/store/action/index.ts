import { PhoneNumber } from '../type'
import axios from 'axios'

export const getPhoneNumber = () => async (dispatch:any) => {

  try {
    const res = await axios.get(`http://localhost:4000/phone-numbers`)
    
    dispatch({
      type: PhoneNumber.GET_PHONE_NUMBERS,
      payload: res.data
    })
  }
  catch (error) {
    dispatch({
      type: PhoneNumber.ERROR,
      payload: error,
    })
  }

}