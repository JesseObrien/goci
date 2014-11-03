var gociApp = angular.module('goci', ['ngRoute', 'goci.controllers', 'goci.services']);

gociApp.config(function($routeProvider, $interpolateProvider) {

	$routeProvider.when('/builds', {
		templateUrl: 'partials/build-list.html',
		controller: 'BuildCtrl'
	})
    .when('/builds/new', {
        templateUrl: 'partials/new-build.html',
        controller: 'BuildCtrl'
    })
	.otherwise({
		redirectTo: '/builds'
	});

	$interpolateProvider.startSymbol('[[');
	$interpolateProvider.endSymbol(']]');
});
