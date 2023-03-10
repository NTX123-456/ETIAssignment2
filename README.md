<h1> Concept of Application: Travel Planning Application </h1>

<h2> Description </h2>
<p> A simple travel planning application that can assist users in better planning their itinerary, allowing them to make more informed decisions about the country or accommodation they are staying in. There are 4 different Microservices, 3 of these microservices have their own seperate database in MYSQL with the exception of the weather update feature utilising a API to display the current weather from the selected country. </p>

<h2> Important Note* </h2>
<p> The code inside this repository has been updated and will not run like before as we have uploaded each of our features on docker and google cloud, do not use these files to test the microservices, another folder will be provided for the microservices, thank you. Please visit this repository instead for microservice testing: https://github.com/NTX123-456/ETIASSIGNMENT2-Microservices- </p>

<h2> Docker and Google Cloud run implementation of our microservices </h2>

<p> We have uploaded each of our 4 features to a docker container and uploaded them to the google cloud run as part of our cloud implementation.This is to ensure that each Microservice would have their own container
 
Here are the google cloud run urls for each of our features, do take note that we currently do not have a front end interface for our features:

1. Itineraries: https://itinerarie-u5obbcs2fa-uc.a.run.app
2. Weather Forecast: https://weatherapp-u5obbcs2fa-uc.a.run.app
3. Viewhotels: https://viewhotels-u5obbcs2fa-uc.a.run.app
4. Booking: https://bookingapp-u5obbcs2fa-uc.a.run.app </p>

<p> 
<h3>Google Cloud Run Service:</h3><br>
 
Itenarie:<br>
![image](https://user-images.githubusercontent.com/73065899/218323495-6ff022e1-f782-431d-927f-6c1f57b11674.png)<br>


Weather Forecast:<br>
![image](https://user-images.githubusercontent.com/73065899/218323394-1f0ee426-0a56-4514-a301-7675f3e53609.png)<br>


Viewhotels:<br>
![image](https://user-images.githubusercontent.com/73065899/218323440-55751207-dc4b-4b1c-8888-d63c26548371.png)<br>


Booking:<br>
![image](https://user-images.githubusercontent.com/73065899/218323457-6dffddb5-1743-46d3-9330-8b15bcdf2432.png)<br>

<h3>Google Cloud Artifact Registry:</h3><br>

Itenarie:<br>
![image](https://user-images.githubusercontent.com/73065899/218323935-bd392f61-6d0c-44d7-bf5c-28e558af4939.png)<br>

Weather Forecast:<br>
![image](https://user-images.githubusercontent.com/73065899/218323966-d5d931ae-f4ed-443d-beb3-208b82b0fd54.png)<br>


Viewhotels:<br>
![image](https://user-images.githubusercontent.com/73065899/218323985-7e86c86b-06f7-44af-80a7-5ddfc5d3ec15.png)<br>


Booking:<br>
![image](https://user-images.githubusercontent.com/73065899/218324019-36938906-662b-4325-b45c-7f72b1e3e389.png)<br>


<h3>Docker Hub:</h3><br>

Weather Forecast and Booking:<br>
![image](https://user-images.githubusercontent.com/73065899/218323732-c54d08d4-217d-457f-80ac-5757b545083e.png)

Itenarie and View Hotels:<br>
![image](https://user-images.githubusercontent.com/73065899/218323748-47b4ad0a-740c-4e26-bb5d-9bad3839d2f9.png)

 

</p>

<h2> How to set up </h2>
<p> 1. For this application there are a total of 3 programs to run (open 3 separate terminals and enter "go run [PROGRAM_NAME]" in their own directory to configure):

1. UserConsole.go (the main console where all the information will be displayed)
2. WeatherForecast.go
3. server.go </p>

<p> 2. Run the 2 database scripts (HotelDatabaseScript.sql & Itineraries DB.sql) in MYSQL to configure.
  
 Database Configuration for HotelDatabase: ![image](https://user-images.githubusercontent.com/73065899/217423817-eadc3fc5-6e7f-434b-bde0-edc1f62b3e3f.png)
<br> Use these credentials to configure the database for hotels: </br>
<br> User: root </br>
<br> Password: Lolman@4567 </br>
<br> connection: 127.0.0.1:3306 </br>
<br>  Database: ETIASSG2_db  </br>
  
 Database Configuration for Itenary database:
  ![image](https://user-images.githubusercontent.com/73065899/217424942-7f38eb76-ec70-4642-a7de-63f238fd9704.png)
<br>Use these credentials to configure the database for hotels:</br>
  <br>User: user </br>
  <br>Password: password </br>
  <br>connection: 127.0.0.1:3306 </br>
  <br> Database: db_itinerarie </br>
  
  For the database for the booking of hotels and attractions feature there are three steps in order to initialize the database:

  1. Create the booking database in MYSQL with these configurations:
   ![image](https://user-images.githubusercontent.com/73065899/217422202-eead80c2-0e1b-4fb8-81ba-e1cedd88a801.png)
    <br>Connection name: booking_db</br>
  2. Run the server file in the booking feature folder: 
  ![image](https://user-images.githubusercontent.com/73065899/217423014-3d92064a-c1a6-43a5-90ab-e0c436b6ad2f.png)
  Once you have seen the message "successfully init db" we can move to the next step
  3. We can jump straight into inserting the data for our table after we have created our booking_db database, do take note that we do not need to manually add our table in the database as it has been automatically added by the server.go file.
  ![image](https://user-images.githubusercontent.com/73065899/217423652-d1773efb-c51e-47b8-ac7d-f33e99db5711.png)


</p>


<h1> Features </h1>

<h2> Planning of itinerarie </h2>
<h3> A feature that allows users to plan out their journey by letting them create, update or view their own personal itineraries. </h3>
<h3> Options </h3>
<h4> 1. View of itinerarie </h4>
<p> This feature is a detailed list of places that users have created into the itinerarie that consists of information like Location, Duration(days), Start Date, and End Date. With all this data results it could be retrieved through MYSQL database with the SELECT statement. </p>

<h4> 2. Create of itinerarie </h4>
<p> This feature will enable users to create an itinerary through a new record in MYSQL database with the INSERT statement, which includes adding of details such as Location, Duration(days), Start Date, and End Date. 

Obtaining users question like:
1. Please enter the location:
2. Please enter the duration of travel (days):
3. Please enter start date (dd/mm/yyyy):
4. Please enter end date (dd/mm/yyyy):
</p>

<h4> 3. Update itinerarie </h4>
<p> This feature is designed to allow users to make changes to the duration of travel, start date, and end date. Subsequently, with the SQL UPDATE statement it will execute the modification towards the exisiting records in the database with the data of Duration(days), Start Date, and End Date. 

Obtaining users question to be updated:
1. Please enter your location to be updated:
2. Please enter the duration of travel to be updated:
3. Please enter the start date to be updated (dd/mm/yyyy):
4. Please enter the end date to be updated (dd/mm/yyyy):
</p>

<h2> Weather Update </h2>
<h3> A feature that allows users to view the current temperature of a specific country or city. </h3>
<h3> Options </h3>
<h4> 1. Displaying of temperature </h4>
<p> This feature will, based on the input country, display in the console the current temperature of that country. Users can also enter the city of that country to obtain a more accurate temperature of a specific location. If the input country or city is not registered, the application will return users to the main menu, printing out the statement "Country not found!" 

The purpose of this feature is to allow users to make a more informed decision on the country they are staying in, and to note whether they need to bring extra layers of clothing for countries that have suddenly gotten colder. </p>

<h2> View Hotel Information </h2>
<h3> A feature that allows users to select and display a list of hotels available in Singapore and other countries </h3>
<h3> Options </h3>
<h4> 1. Listing All Hotels </h4>
<p> This feature allows the user to display all the hotels available in the database with no filters. </p>

<h4> 2. Display Hotels from a certain country </h4>
<p> This feature allows the user to input a certain country to display a hotel from. After the user has input a country they want, the console will display all the hotels from that particular country. 
  
Current countries available in the database:
1. Singapore
2. Malaysia
3. Thailand
4. Indonesia
5. Australia
</p>

<h4> 3. Display Hotels with a certain rating </h4>
<p> This feature allows the user to display all the hotels with a certain rating. For example, the user can input the number 4 into the console and it will retrieve all the 4 star hotels from the database and display it. </p>

<h4> 4. Display Hotels with certain amenities </h4>
<p> This feature allows the user to display all the hotels with certain amenities. For example, if the user wants to view hotels with Pools, the user can input pools and the console will retrieve all the hotels with pools in their amenities. </p>

<h4> 5. Display Hotels from a certain price range </h4>
<p> This feature allows the user to display hotels from a certain price range they want. There are a total of 5 price ranges the user can choose from

  1. Hotels below $50
  2. Hotels between $50 and $100
  3. Hotels between $100 and $200
  4. Hotels between $200 and $300
  5. Hotels between $300 and $400
</p>




<h2> Booking Hotels & Attractions </h2>
<h3> A feature that allows user to view a list of hotel and attraction prices, after which user can start to book them. User can also choose to update or cancel their booking. </h3>

<h3> Options </h3>
<h4> Guide to start off booking feature:
<h4> User will be required to input their name to begin with this feature. </h4>
<h4> 1. Listing all tourist attractions and hotels </h4>
<p> This option allow user to list out all tourist attractions and hotels that are available with the pricing. </p>

<h4> 2. Book tourist attractions and hotels </h4>
<p> This option allows user to book an attraction or hotel by choosing the correct id, once id is confirmed, click on enter and the booking will be successful. </p>
<p> Example of Attractions and Hotels are in Singapore???s context. </p>

<h4> 3. Cancel or update booking </h4>
<p> This option allows user to either cancel a booking by typing out ???cancel <id>??? or update a booking by typing out ???update <id>???, where id is the number located on the left of each attraction or hotel. 
For update of booking the console will prompt user to pick another attraction or hotel by keying in the ideal id, after which the booking will be updated successfully.
</p>

<h4> Navigation options </h4>
<p> Type ???main??? to return to main menu. </p>
<p> Type ???quit??? to exit the program. </p>
