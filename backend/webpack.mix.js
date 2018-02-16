const {mix} = require('laravel-mix');

/*
 |--------------------------------------------------------------------------
 | Mix Asset Management
 |--------------------------------------------------------------------------
 |
 | Mix provides a clean, fluent API for defining some Webpack build steps
 | for your Laravel application. By default, we are compiling the Sass
 | file for the application as well as bundling up all the JS files.
 |
 */


mix.copy([
    'resources/assets/vendor'
], 'static/vendor/');
mix.copy([
    'resources/assets/img'
], 'static/img/');

mix.sass('resources/assets/sass/app.sass', 'static/css');

mix.combine(['resources/assets/js/*'], 'static/js/app.js');

// mix.browserSync('localhost:8080');