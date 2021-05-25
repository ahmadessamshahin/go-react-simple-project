type Column = {
    id: 'number' | 'countryCode' | 'country' | 'state' | 'action';
    label: string;
    minWidth: number;
}

const columns: Column[] = [
    {
        id: 'number',
        label: 'Number',
        minWidth: 200
    },
    {
        id: 'countryCode',
        label: 'Country Code',
        minWidth: 100
    },
    {
        id: 'country',
        label: 'Country',
        minWidth: 170
    },
    {
        id: 'state',
        label: 'State',
        minWidth: 170
    },
    {
        id: 'action',
        label: 'Actions',
        minWidth: 100
    }
];

export default columns;
