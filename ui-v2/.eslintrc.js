const path = require('path');

module.exports = {
  extends: ['react-app', 'airbnb'],
  rules: {
    'react/jsx-filename-extension': [1, { 'extensions': ['.js', '.jsx'] }],
  },
  settings: {
    'import/resolver': {
      node: {
        paths: [path.resolve(__dirname, 'src')],
      }
    }
  }
}