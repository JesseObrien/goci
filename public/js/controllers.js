var gociControllers = angular.module('goci.controllers', []);

gociControllers.controller('BuildCtrl', function($scope, $rootScope, Build) {

	$rootScope.breadCrumbs = [
		{ label: "Builds", url: "#/builds" }
	];

	$scope.newBuild = {};

	$scope.builds = Build.query();

	$scope.createNewBuild = function() {
		var nb = $scope.newBuild;
		var b = new Build();

		b.name = nb.name;
		b.description = nb.description;
		b.vc_type = nb.vc_type;
		b.vc_url = nb.vc_url;

		b.$save();
		
		$scope.newBuild = {};
	};

});
