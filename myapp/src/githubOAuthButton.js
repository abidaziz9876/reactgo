

import React from 'react';

const GitHubOAuthButton = () => {
  const handleInstall = () => {
    // Redirect directly to GitHub OAuth URL
    window.location.href = 'https://github.com/apps/MyApp9876/installations/select_target';
  };

  return (
    <div className="GitHubOAuthButton">
      <button onClick={handleInstall}>Install GitHub App</button>
    </div>
  );
};

export default GitHubOAuthButton;
