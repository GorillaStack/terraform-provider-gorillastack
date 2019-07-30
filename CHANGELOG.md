## 0.2.1 (July 31, 2019)

#### BREAKING CHANGES:
#### FEATURES:
#### IMPROVEMENTS:
#### BUG FIXES:

* Fixing an infinite loop when stringifying the StringArrayOrNull type in the context

#### NOTES:

* Version bumped to version 0.2.1 for this bugfixj


## 0.2.0 (July 26, 2019)

#### BREAKING CHANGES:

* CloudTrail triggers must specify match_fields as a single list item rather than a map. There was an oversight regarding the values of the map being arrays of strings, while Terraform maps only support strings

#### FEATURES:


* **New Trigger Type for Rule Resources:** CloudTrail Event
* **New Action Type for Rule Resources:** CheckTagCompliance
* **New Action Type for Rule Resources:** NotifyEvent (for CloudWatch Events)

#### IMPROVEMENTS:

* Improving toggling of adding the API version to the path so this doesn't interfere with local development when the API destination is overriden.

#### BUG FIXES:

* Fixing broken implementation of the CloudTrail trigger action (see notes in breaking changes above)

#### NOTES:

* Version bumped to version 0.2.0 as there are some minor breaking changes and significant new features, thus a patch version is not quite appropriate.


## 0.1.0 (July 8, 2019)

#### BREAKING CHANGES:

* None

#### FEATURES:

* **New Resource:** Rule
* **New Resource:** Tag Group

#### IMPROVEMENTS:
#### BUG FIXES:
#### NOTES:

* This is the initial GorillaStack Terraform provider implementation. As implied by a < 1 major version, there is little commitment to maintaining backwards compatibility right now, and there may be breaking changes in coming minor version updates. All breaking changes will be communicated in the changelog.
