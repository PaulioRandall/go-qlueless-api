# Application Decision Register

All notable decisions made within this project will be documented here. I've decided to use a simple markdown structure to maintain consistency with the API changelog, and promote human and machine readability.

Since my experimentation of ADRs is starting part way through the project many entries you may expect to see won't be listed. I will attempt to back populate as I come across through cleaning and refactoring sessions. Depending on how they are used and the actual content they end up containing, in future I will reconsider whether the changelog and ADR should be separate documents with separate concerns or amalgamated for ease of user reference. The decisions may also need to be split based on whether they refer to the API or the source code.

I decided to introduce an Application Decision Register (ADR) as I've always wanted to learn or create an effective, yet lightweight, method for recording key decisions. I've always worried that the making of decisions within software engineering have mostly been based on personal preferences and intuitions rather than proven means such as anchoring on a data derived baseline and adapting to fit the context at hand.

For complex junctions, the prior results in sub-optimal decisions that are often the cause for project delays, technical debt, cat fights, and a notable reduction in team performance and member engagement. Of course good data is often hard to come by and an appropriate baseline needs some agreement but  I'm supporting the idea that the investment is worth it.

## Usage of set based API endpoints

The experimental use of set based API endpoints is an attempt to simplify the interface. My hypothesis is that by using only sets of resources within requests and responses:

- the number of endpoints can be reduced thus simplifying it.
- the emphasize is placed on solving the harder problem of dealing with multiple resources from the outset.
- both front-end and back-end developers only need think about multiple resources during implementing, if they choose to.

Notable issues and disadvantages include:

- the handling of content types which can only represent single resources.
  - In this case only JSON is supported which handles multiple resources using arrays or object maps.
- the confusion that may result due to the diversion from the familiar `/set` with `/set/{id}` approach.
  - Any experimental deviation from the norm will face this issue. The fact that the deviation is small combined with good documentation should alleviate this issue.