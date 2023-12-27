import React from 'react';
import GitHubOAuthButton from './githubOAuthButton';

const App = () => {
  return (
    <div className="App">
      <header className="App-header">
        <p>
          Welcome to the GitHub OAuth App! Click the button below to log in with GitHub.
        </p>
        {/* Render the GitHub OAuth button component */}
        <GitHubOAuthButton />
      </header>
    </div>
  );
};

export default App;
