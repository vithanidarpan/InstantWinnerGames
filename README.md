# Big Picture
Big picture is in an other document but to make it short :
- InstantWinnerGames can be created and linked to Places. Meaning they can only be played if you are close enough to the Place of the InstantWinnerGame.
InstantWinnerGames have a start/endDate : they cannot be playout outside this time range. InstantWinnerGames have also a Won status (true/false)

- Players can request for InstantWinnerGames they can play, based on their GPS and time (server time)

- A player can only play once to an InstantWinnerGame (we provide a fingerpring, outside of this project). Everytime a player plays to an InstantWinnerGame, we create a InstantWinnerPlayer instance, for stats purposes.

- The player who wons is the one that plays the first after InstantWinnerGames.playTime. You can see details below


# Install

go get .. whatever will be requested

go build -o service-api

./service-api


# Need to be done
Switch to mysql

Add small populate go script with a few examples

Add 1 apis : GetListOfOpenedInstantGames 

Should take as input latitude and longitude

Should return all InstantWinnerGames + Place where

InstantWinnerGames.won == false && InstantWinnerGames.startDate< serverTime < InstantWinnerGames.endDate 

&& Place.id==InstantWinnerGames.id && (latitude-0.1)< Place.latitude <(latitude+0.1) && (longitude-0.1)< Place.longitude <(longitude+0.1)


Modify CreateInstantWinnerPlayer

Should take as input a InstantWinnerGames id, latitude and longitude and returns according to below procedure 

(I have a complex issue there. I could have concurrent calls ? Let’s say we will have one single server is enough ?)

Here is detailed process.

First we get server time => serverTime. 

Then we get matching InstantWinnerGame
And

0/ We check if InstantWinnerGame is not won
If InstantWinnerGame.won => returns{
    result : false,
    code : 1,
    message : 'Already won game'
}

1/ We check if InstantWinnerGame is linked to a Place whose latitude and longitude are OK (see above API)

If GPS is not OK returns{
    result : false,
    code : 2,
    message : 'Not allowed from here'
}

else (We check if InstantWinnerGame is still available)

If !(InstantWinnerGame.startTime < serverTime < InstantWinnerGame.endTime) => returns{
    result : false,
    code : 2,
    message : 'game is over'
}

From then on, we know InstantWinnerGame is open and available for the player

2/ We check if this player already played in this game

check if one InstantWinnerPlayer where 
InstantWinnerPlayer.InstantWinnerGameID==instantWinnerGameId &&
(InstantWinnerPlayer.fp==fp || InstantWinnerPlayer.email==email)

If true returns {
    result : false,
    code : 3,
    message : 'Already played'
} 

else

2/ It means that game is open, player allowed to play

if (serverTime>InstantWinnerGame.playTime && InstantWinnerGame.won == false) then

a/ set InstantWinnerGame.won=true

b/ save new InstantWinnerPlayer with result=true and time=serverTime

c/ returns{
    result : true,
    message : 'game won'
}

else (played but too early)

a/ save new InstantWinnerPlayer with result=false and time=serverTime

b/ returns{
    result : false,
    code : 4,
    message : 'you lost'
}




Dockerfile and Dockerfile composer
