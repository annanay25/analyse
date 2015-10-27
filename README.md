# analyse  
Prediction model based on Collaborative Filtering.  

Implementation based on Implicit feedback as explained mathematically in http://yifanhu.net/PUB/cf.pdf  

What this package does:  
1. Connects to a shop in Reaction Commerce and gets all shop data from the Cassandra database.  
2. Listens to all the view and buy events.  
3. Implements ALS through Collaborative filtering using appropriate data modeling.  
4. Returns an array of recommendations to the user.  

Functions and API:  
1. Event listener which listens to view and buy events.  
2. Exposed recommender API endpoint which gives list of recommendations.  

Routes:  
1. POST /event  
2. GET /query  
3. POST /AddUser
4. POST /AddItem

Usage:  

<!-- Will be updated soon -->
