import { combineReducers } from 'redux'
import PhoneNumberReducer from './phoneReducer'
export default combineReducers({
  phoneNumberList: PhoneNumberReducer
})