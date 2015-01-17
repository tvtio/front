var gulp = require('gulp');
var jshint = require('gulp-jshint');

gulp.task('jshint', function () {
    gulp.src('*.js')
        .pipe(jshint())
        .pipe(jshint.reporter('default'));
});

gulp.task('default', [], function() {});