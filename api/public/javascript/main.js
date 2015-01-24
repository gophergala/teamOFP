var Groupify = angular.module('Groupify', []);

Groupify.controller('MainCtrl', function($scope, $http) {
  $scope.queue = [];
  $scope.query = "";

  $http.get('/api/v1/queue/list')
  .then(function(res){
    $scope.queue = res.data;
  });

  $scope.search = function(){
    // TODO bring spotify search logic here
    console.log($scope.query);
  }
});
