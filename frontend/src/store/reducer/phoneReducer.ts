import { PhoneNumber } from '../type'

const initialState = {
  phoneNumber: [],
  loading: true
}

export default function reducer(state = initialState, action: any) {

  switch (action.type) {

    case PhoneNumber.GET_PHONE_NUMBERS:
      return {
        ...state,
        users: action.payload,
        loading: false
      }
    case PhoneNumber.ERROR:
      return {
        loading: false,
        error: action.payload
      }

    default: return state
  }

}