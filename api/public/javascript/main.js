var Groupify = angular.module('Groupify', []);

Groupify.controller('MainCtrl', function($scope, $http) {
  var spotifyApi = new SpotifyWebApi();
  $scope.queue = [];
  $scope.query = "";
  $scope.trackResults = [];

  $http.get('/api/v1/queue/list')
  .then(function(res){
    $scope.queue = res.data;
  });

  $scope.search = function(){
    spotifyApi.searchTracks($scope.query, {limit: 10, offset: 0}, function(err, data) {
      $scope.trackResults = data.tracks.items;
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
