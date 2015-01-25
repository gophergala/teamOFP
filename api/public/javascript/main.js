var Groupify = angular.module('Groupify', []);

Groupify.controller('MainCtrl', function($scope, $http, $timeout) {
  var spotifyApi = new SpotifyWebApi();
  $scope.queue = [];
  $scope.query = "";
  $scope.trackResults = [];


  (function tick() {
    $http.get('/api/v1/queue/list')
    .then(function(res){
      $scope.queue = res.data;
      $timeout(tick, 1000);
    });
  })();

  $scope.search = function(){
    spotifyApi.searchTracks($scope.query, {limit: 10, offset: 0}, function(err, data) {
      $scope.trackResults = data.tracks.items;
    });
  };

  $scope.dequeue = function(track){
    $http.post('/api/v1/queue/delete', {
      track_id: track.id
    })
    .then(function(res){
      // FIXME: handle errors
      console.log("de-queued track " + track.name);
    });
  };

  $scope.enqueue = function(track){
    $http.post('/api/v1/queue/add', {
      track_id: track.id
    })
    .then(function(res){
      console.log("Enqueued track " + track.name);
    });
  };

});
