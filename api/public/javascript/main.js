var Groupify = angular.module('Groupify', ['timer']);

Groupify.controller('MainCtrl', function($scope, $http, $timeout) {
  var spotifyApi = new SpotifyWebApi();
  $scope.queue = [];
  $scope.query = "";
  $scope.trackResults = [];
  $scope.current_track = { artist: "Marisa Monte", name: "Barulhinho Bom", time_remaiming: 97 };

  (function tick() {
    $http.get('/api/v1/queue/list')
    .then(function(res){

      // sum time to play up to each track in collection
      var sum = 0;

      if ($scope.current_track) { sum += parseInt($scope.current_track.time_remaiming); }

      for(var i = 0; i < res.data.length; i++) {
        track = res.data[i];
        sum += parseInt(track.time);
        track.time_to_play = sum;
      }

      $scope.queue = res.data.queue;
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

Groupify.filter('secondsToTime', function() {
  // shameless copy/paste
  // http://codeaid.net/javascript/convert-seconds-to-hours-minutes-and-seconds-(javascript)
  return function(secs) {
    var hr  = Math.floor(secs / 3600);
    var min = Math.floor((secs - (hr * 3600))/60);
    var sec = secs - (hr * 3600) - (min * 60);

    if (hr  < 10) { hr  = "0" + hr; }
    if (min < 10) { min = "0" + min;}
    if (sec < 10) { sec = "0" + sec;}
    if (hr)       { hr  = "00"; }
    return hr + ':' + min + ':' + sec;
  };
});

