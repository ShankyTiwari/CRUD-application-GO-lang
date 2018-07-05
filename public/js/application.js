$(document).ready(function() {
  var end_of_track = 15;
  var players = [
    {
      id: 1,
      position: 0,
      keycode: 80,
      css_locator: '#player1_strip'
    },
    {
      id: 2,
      position: 0,
      keycode: 81,
      css_locator: '#player2_strip'
    }
  ];

  function advancePlayer(keyCode) {
    players.forEach(function(player){
      if (keyCode == player.keycode && players[0].position < end_of_track && players[1].position < end_of_track) {
        move_player(player.css_locator);
        player.position++;
        winner(player.position, player.id);
      }
    });
  }

  $(document).keyup(function(e) {
    advancePlayer(e.keyCode);
  });
});

function move_player(strip) {
  $(strip).find('td.active').removeClass('active').next().addClass('active');
}

function winner(player, num) {
  if (player > 14) {
    alert("Player " + num + " has won!");
  }
}