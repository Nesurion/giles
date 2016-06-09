function archivMovieRequest(title, dirname){
	$.ajax({
	    url: 'api/archiv/movie',
	    type: 'POST',
	    data: {title: title, dirname: dirname}
	});
}

function archivShowRequest(title, dirname){
	$.ajax({
	    url: 'api/archiv/show',
	    type: 'POST',
	    data: {title: title, dirname: dirname}
	});
}

function deleteRequest(title, dirname){
	$.ajax({
	    url: 'api/delete',
	    type: 'POST',
	    data: {title: title, dirname: dirname}
	});
}
