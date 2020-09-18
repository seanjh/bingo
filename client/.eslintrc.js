module.exports = {
  env: {
    browser: true,
    es2021: true,
  },
  extends: ['eslint:recommended', 'airbnb-base'],
  ignorePatterns: ['dist/', 'node_modules'],
  parser: '@typescript-eslint/parser',
  parserOptions: {
    ecmaVersion: 12,
    sourceType: 'module',
  },
  plugins: ['@typescript-eslint'],
  rules: {
    'array-callback-return': ['error', { checkForEach: true }],
    complexity: ['error', { max: 10 }],
    eqeqeq: ['error', 'always'],
    'no-eval': 'error',
  },
};
