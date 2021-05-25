import { Box, Container } from '@material-ui/core';
import './App.css';
import List from './components/list/List';
import Navbar from './components/Navbar';
import Router from './components/Router';

function App() {
    return (
        <div className="App">
            <Navbar />
            <Container>
                <Box m={4}>
                    <Router />
                </Box>
            </Container>
        </div>
    );
}

export default App;
