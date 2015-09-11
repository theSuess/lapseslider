$('#testButton').click(function(){
	$.get("/blink",function(data){
		if (data !== 'OK'){
			alert('The test returned an error. Please checkout the Logs');
		}
	});
});
