postgres:
	docker run --name image-container -e POSTGRES_PASSWORD=milan123 -e POSTGRES_USER=imageInfo -e POSTGRES_DB=image -p 5436:5432 -d postgres

createdb:
	docker exec -it image-container createdb --username=imageInfo --owner=imageInfo newdb 

dropdb:
	docker exec -it bank dropdb --username=milanbank milanbank