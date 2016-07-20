jQuery(document).ready(function($){
	$("#sendMessageForm").submit(function(e){
		$.ajax({
			ur: "/api/messages"
		})
			.done(function(response, textStatus, xhr){
				//console.log("Done", arguments);
				if (xhr.status == 200) {
					//location.reload();
				} else {
					console.warn(arguments);
				}
			}).fail(function(xhr, textStatus, error){
				console.error(arguments);
			});
		e.preventDefault();
	});
});