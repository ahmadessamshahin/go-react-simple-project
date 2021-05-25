enum PhoneState {
    Valid = 'Valid',
    Invalid = 'Invalid'
}

type PhoneNumber = {
    id: number;
    number: string;
    countryCode: string;
    country: string;
    state: PhoneState;
}

const phoneNumbers: PhoneNumber[] = [
    { id: 1, number: '(20) 1091054748', countryCode: '20', country: 'Egypt', state: PhoneState.Valid },
    { id: 2, number: '(20) 10910547480', countryCode: '20', country: 'Egypt', state: PhoneState.Invalid }
];

export default phoneNumbers;
