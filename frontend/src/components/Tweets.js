import React from "react";
import { Route } from "react-router-dom";
import Tweet from "./Tweet.js";
import TweetList from './TweetList.js'

const Tweets = () => {
	return (
		<div>
			<Route exact path="/" component={TweetList} />
			<Route path={`/tweets/:id`} component={Tweet} />
		</div>
	);
};


export default Tweets;
