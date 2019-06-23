pragma solidity >=0.4.22 <0.6.0;

contract Mesh {
    enum UserType {
        None, Sub, Pub//, PubSub
    }
    
    struct Users {
        UserType userType;
        mapping(string => bool) topicList;
    }
    
    mapping(address => Users) userMap;
    
    function subscribe(address a, string memory topic) public {
        userMap[a].userType = UserType.Sub;
        userMap[a].topicList[topic] = true;
    }
    
    function publish(address a, string memory topic) public {
        userMap[a].userType = UserType.Pub;
        userMap[a].topicList[topic] = true;
    }
    
    function isPeerAPublisher(address a, string memory topic) public view returns (bool) {
        return (userMap[a].userType == UserType.Pub && userMap[a].topicList[topic] == true);
    }
    
    function isPeerSubscribed(address a, string memory topic) public view returns (bool) {
           return (userMap[a].userType == UserType.Sub && userMap[a].topicList[topic] == true);
    }
}