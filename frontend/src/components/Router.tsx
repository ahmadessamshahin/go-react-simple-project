import React from 'react';
import { Route, Switch, BrowserRouter } from 'react-router-dom';
import List from './list/List';
import View from './view/View';

class Router extends React.Component {
	render() {
		return (
			<BrowserRouter>
				<Switch>
					{/* <Route path='/:phoneNumberId/edit' component={Edit} /> */}
                    <Route path='/:phoneNumberId' component={View} />
					<Route path='/' component={List} />
				</Switch>
			</BrowserRouter>
		);
	}
}

export default Router;
