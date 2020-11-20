# GiftsWinner
go get .. whatever needed
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
Should take as input a InstantWinnerGames id and returns
I have a complex issue there. I could have concurrent calls ? Let’s say we will have one single server is enough ?

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
else (We check if InstantWinnerGame is still available)
If !(InstantWinnerGame.startTime < serverTime < InstantWinnerGame.endTime) => returns{
    result : false,
    code : 2,
    message : 'game is over'
}

From then on, we know InstantWinnerGame is open
1/ We check if this player already played in this game
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
