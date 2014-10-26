var gociServices = angular.module('goci.services', ['ngResource']);

gociServices.factory('Build', function($resource) {
	return $resource('/builds/:buildId', { buildId: '@id' });
});
