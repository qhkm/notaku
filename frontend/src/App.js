import React, { Component } from "react";
import { BrowserRouter, Route } from "react-router-dom";
import "./App.css";
import Navbar from "./components/Navbar.js";
import Tweets from "./components/Tweets.js";
import Notifications from "./components/Notifications.js";
import Explore from "./components/Explore.js";
import Messages from "./components/Messages.js";
import Tweet from "./components/Tweets.js";

class App extends Component {
	render() {
		return (
			<BrowserRouter>
				<div>
					<Navbar />
					<Route exact path="/" component={Tweets} />
					<Route exact path="/explore" component={Explore} />
					<Route
						exact
						path="/notifications"
						component={Notifications}
					/>
					<Route exact path="/messages" component={Messages} />
					<Route path={`/tweets/:id`} component={Tweet} />
				</div>
			</BrowserRouter>
		);
	}
}

export default App;
