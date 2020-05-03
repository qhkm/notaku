import React  from "react";

const Tweet = props => {
	const {tweet} = props.location.state
	// var tweet = props.tweet.find(tweet => tweet.id === props.match.params.id);

	return (
		<div className="container">
			<div className="space" />
			<main className="content-area">
				<div className="card">
					{tweet}
				</div>
			</main>
		</div>
	);
};

export default Tweet;
