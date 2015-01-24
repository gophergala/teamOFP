var Groupify = angular.module('Groupify', []);

Groupify.controller('MainCtrl', function($scope, $http) {
  $scope.queue = [];

  $http.get('/api/v1/queue/list')
  .then(function(res){
    $scope.queue = res.data;
  });
});
