module.exports = {
  branches: ["master"],
  ci: false,
  dryRun: false,
  plugins: [
    "@semantic-release/commit-analyzer",
    "@semantic-release/release-notes-generator",
    "@semantic-release/github",
  ],
};
