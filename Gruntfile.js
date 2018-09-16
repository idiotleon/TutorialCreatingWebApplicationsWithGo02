const sass = require('node-sass');

module.exports = function (grunt) {
    require('load-grunt-tasks')(grunt);

    grunt.initConfig({
        sass: {
            options: {
                implementation: sass,
                sourceMap: true
            },
            dist: {
                files: {
                    'public/css/app.css': 'sass/app.scss'
                }
            }
        }
    });

    grunt.registerTask('default', ['sass']);
};
