var app=require('express')();

var http=require('http').Server(app);

var io = require('socket.io')(http);


io.on('connection', function(socket){
    socket.on('chat message', function(msg){

      io.emit('chat message',msg);
      console.log('message: ' + msg);
    });

    socket.broadcast.emit('hi');
});


app.get('/',function(req,res){
    res.send('<h1>Express Socket</h1>')
})

app.get('/chat',function(req,res){
    res.sendFile(__dirname+'/index.html');
})

http.listen(3322,function(){
    console.log("listening on *:3322")
})
