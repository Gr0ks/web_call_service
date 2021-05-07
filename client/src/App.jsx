import React, { useState } from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";

function App() {
    return <div className="App">
		<BrowserRouter>
			<Switch>
				<Route path="/" exact></Route>
				<Route path="/room/:roomID"></Route>
			</Switch>
		</BrowserRouter>
	</div>;
}

export default App;