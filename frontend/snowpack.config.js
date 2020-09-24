module.exports = {
  installOptions: {
    installTypes: true,
  },
  mount: {
    public: '/',
    src: '/_dist_',
  },
  plugins: [
    [
      '@snowpack/plugin-run-script',
      {
        cmd: 'tsc --noEmit',
        watch: '$1 --watch',
      },
    ],
    [
      '@snowpack/plugin-run-script',
      {
        cmd: "eslint 'src/**/*.{js,jsx,ts,tsx}'",
        watch: 'watch "$1" src',
      },
    ],
  ],
};
