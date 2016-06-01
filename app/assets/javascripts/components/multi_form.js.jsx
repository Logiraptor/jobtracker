class MultiForm extends React.Component {
	constructor(props) {
		super(props);
		var values = props.values;
		values.push('');
		this.state = {
			values: values,
			focused: 0
		}
	}

	componentDidMount() {
		this.compact();
	}

	componentDidUpdate(prevProps, prevState) {
	  this.compact();  
	}

	onKeyDown(event) {
		if (event.keyCode === 13) {
			console.log(event.target.dataset.index);
			event.preventDefault();
			var newValues = this.state.values;
			newValues.push('');
			this.setState({
				values: newValues,
				focused: +(event.target.dataset.index) + 1
			})
		}
	}

	onChange(event) {
		var newValues = this.state.values;
		newValues[event.target.dataset.index] = event.target.value;
		this.setState({
			values: newValues
		});
	}

	onFocus(event) {
		this.setState({
			focused: event.target.dataset.index
		})
	}

	compact() {
		var hasBlank = false;
		var newValues = [];
		this.state.values.forEach((value, i) => {
			if (!value && i == this.state.focused) {
				hasBlank = true;	
				newValues.push(value);
			}
			if (value) {
				newValues.push(value);
			}
		});
		
		if (newValues.length != this.state.values.length) {
			this.setState({
				values: newValues
			})
		}
	}

	render() {
		var forms = this.state.values.map((value, i) => {
			return (
				<input key={i}
					autoFocus
					data-index={i}
					type="text"
					value={value}
					name={this.props.name}
					onChange={this.onChange.bind(this)}
					onKeyDown={this.onKeyDown.bind(this)}
					onFocus={this.onFocus.bind(this)}
					{...this.props.attrs}/>
			)
		})

		return (
			<div>
				{forms}
			</div>
		);
	}
}