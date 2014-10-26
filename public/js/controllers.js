var gociControllers = angular.module('goci.controllers', []);

gociControllers.controller('BuildCtrl', function($scope, $rootScope, Build) {

	$rootScope.breadCrumbs= [
		{ label: "Builds", url: "#/builds" }
	];

	$scope.builds = Build.query();

});
